package router

import (
	"redis-learn/controller"
	"redis-learn/middleware"

	"github.com/gin-gonic/gin"
)

func Start() {
	//gob.Register(dto.UserDTO{})
	r := gin.Default()
	//store := cookie.NewStore([]byte("secret42353"))
	//r.Use(sessions.Sessions("mysession", store))
	r.Use(middleware.RefreshToken())

	user := r.Group("/user")
	{
		user.POST("/code", controller.SendCode)
		user.POST("/login", controller.Login)
		user.GET("/me", middleware.Filter(), controller.Me)
		// TODO /info/:id
		// TODO /:id
		// TODO /user/sign
		/*
			1. 获取当前登录用户的日期
			2. 获取日期
			3. 拼接key
			4. 获取今天是本月的第几天
			5. 写入Redis SETBIT key offset 1
		*/
		// TODO  /user/sign/count
	}

	follow := r.Group("/follow", middleware.Filter())
	{
		follow.PUT("/:id/:isFollow", controller.FollowController)
		follow.GET("/or/not/:id", controller.IsFollow)
		//TODO /common/:id
	}
	shopType := r.Group("/shop-type")
	{
		shopType.GET("/list", controller.QueryShopList)
	}

	blog := r.Group("/blog")
	{
		blog.GET("/hot", controller.QueryHotBlog)
		blog.GET("/:id", controller.QueryBlogById)
		blog.PUT("/like/:id", controller.BlogLike)
		blog.GET("/likes/:id", controller.BlogLikes)
		//TODO POST 保存blog,并实现推送
		/* 见service
		 */
		//TODO GET:/blog/of/follow lastId:上一次查询的最小值,offset:偏移量
		/*
			1.获取当前用户
			2.查询收件箱
			3.解析blogId,minTime,offset
			4.根据id查询blog
			5.封装并返回

		*/
	}

	shop := r.Group("/shop")
	{
		shop.POST("", controller.SaveShop)
		shop.PUT("", controller.UpdateShop)
		shop.GET("/:id", controller.QueryShopById)
		//TODO GET:/shop/of/name?name=
		//TODO GET /shop/of/type?typeId= & current= &x= &y=
		/*
			1.判断是否需要根据坐标查询
			2.计算分页参数
			3.查询redis 按照距离排序,分页.结果:shopId,distance
			4.解析id
			5.根据id查询shop
			6.返回

		*/
		/*
			把店铺放到redis:
			1. 查询店铺信息
			2. 把店铺分组,按照typeId分组,typeId一致的放到一个集合
			3. 分批完成写入redis
				3.1 获取类型id
				3.2 获取同类型的店铺的集合
				3.3 写入redis GEOADD key 经度 维度 member
		*/
	}

	voucher := r.Group("/voucher")
	{
		voucher.POST("/seckill", controller.AddSeckillVoucher)
		voucher.POST("", controller.AddVoucher)
		voucher.GET("/list/:shopId", controller.QueryVoucherOfShop)
	}
	voucherOrder := r.Group("/voucher-order")
	{
		voucherOrder.POST("/seckill/:id", middleware.Filter(), controller.SeckillVoucher)
	}

	r.Run(":8081")
}
