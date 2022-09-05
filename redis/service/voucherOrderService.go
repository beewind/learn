package service

import (
	"redis-learn/dao"
	"redis-learn/dto"
	"redis-learn/entity"
	"redis-learn/utils"
	"strconv"
	"time"
)

type VoucherOrderService struct{}

/*
	版本二
	使用redis的锁解决了高并发下的一人一单:
	存在问题:
		1.不可重入
		2.不可重试
		3.超时释放
		4.主从一致
		5.直接操作在数据库,吞吐量低
		6.锁线程不安全
*/
func (v *VoucherOrderService) SeckillVoucher(voucherId int, userId int) dto.Result {
	var result dto.Result
	// 1.查询优惠卷
	voucher, err := dao.SelectVoucherSeckillId(voucherId)
	if err != nil { // 我的数据库延迟比较高,高并发时会:Error 1040: Too many connections
		return result.Fail("查询失败:" + err.Error())
	}
	// 2.判断秒杀是否开始
	if time.Time(voucher.BeginTime).After(time.Now()) {
		return result.Fail("活动还未开始")
	}

	// 3.判断秒杀是否结束
	if time.Now().After(time.Time(voucher.EndTime)) {
		return result.Fail("活动已结束")
	}

	// 4.判断库存是否充足
	if voucher.Stock <= 0 {
		return result.Fail("已售罄")
	}

	// 5.扣减库存
	//		一人一单
	// 		5.1 查询订单
	lock := "order:" + strconv.Itoa(userId)
	isLock, _ := utils.TryLock(lock)

	if !isLock {
		return result.Fail("不允许重复下单!")
	}
	//		5.2 判断是否存在
	count := dao.SelectWithUserIdAndVoucherId(userId, voucherId)
	if count > 0 {
		// 			用户已经购买过了
		return result.Fail("用户已购买一次")
	}

	err = dao.SubSeckillVoucher(voucherId)
	if err != nil {
		return result.Fail("秒杀失败:" + err.Error())
	}

	// 6.创建订单
	// 6.1 订单id
	var voucherOrder entity.VoucherOrder
	var redisIdWork utils.RedisIdWorker
	voucherOrderId := int(redisIdWork.NextId("order"))
	voucherOrder.Id = voucherOrderId

	// 6.2 用户id
	voucherOrder.UserId = userId

	// 6.3 代金券id
	voucherOrder.VoucherId = voucherId

	err = dao.CreateVoucherOrder(voucherOrder)
	if err != nil {
		result.Fail("订单创建失败")
	}
	utils.Unlock(lock)
	// 7.返回订单id
	return result.Ok(voucherOrderId)
}

/*
	版本一
	使用乐观锁,解决了超卖问题
	问题:用户可以重复下单
*/
/* func (v *VoucherOrderService) SeckillVoucher(voucherId int, userId int) dto.Result {
	var result dto.Result
	// 1.查询优惠卷
	voucher, err := dao.SelectVoucherSeckillId(voucherId)
	if err != nil { // 我的数据库延迟比较高,高并发时会:Error 1040: Too many connections
		return result.Fail("查询失败:" + err.Error())
	}
	// 2.判断秒杀是否开始
	if time.Time(voucher.BeginTime).After(time.Now()) {
		return result.Fail("活动还未开始")
	}

	// 3.判断秒杀是否结束
	if time.Now().After(time.Time(voucher.EndTime)) {
		return result.Fail("活动已结束")
	}

	// 4.判断库存是否充足
	if voucher.Stock <= 0 {
		return result.Fail("已售罄")
	}

	// 5.扣减库存
	//		一人一单
	// 		5.1 查询订单
	_, err = dao.SelectWithUserIdAndVoucherId(userId, voucherId)

	//		5.2 判断是否存在
	if err == nil {
		// 			用户已经购买过了
		return result.Fail("无法重复下单!")
	}

	err = dao.SubSeckillVoucher(voucherId)
	if err != nil {
		return result.Fail("秒杀失败:" + err.Error())
	}

	// 6.创建订单
	// 6.1 订单id
	var voucherOrder entity.VoucherOrder
	var redisIdWork utils.RedisIdWorker
	voucherOrderId := int(redisIdWork.NextId("order"))
	voucherOrder.Id = voucherOrderId

	// 6.2 用户id
	voucherOrder.UserId = userId

	// 6.3 代金券id
	voucherOrder.VoucherId = voucherId

	err = dao.CreateVoucherOrder(voucherOrder)
	if err != nil {
		result.Fail("订单创建失败")
	}

	// 7.返回订单id
	return result.Ok(voucherOrderId)
}
*/
