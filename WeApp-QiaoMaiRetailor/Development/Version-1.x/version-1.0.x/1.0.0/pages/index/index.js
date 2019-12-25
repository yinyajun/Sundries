//index.js
//获取应用实例
const app = getApp()

Page({
  //页面初始数据，注意要用绝对路径
  data: {
    imgUrls: [
      '/images/swiper_01.png',
      '/images/swiper_02.jpg',
      '/images/swiper_03.png',
    ],
    indicatorDots: true,
    autoplay: true,
    interval: 5000,
    duration: 1000,
    proList: [
      {
        logo: '/images/index_item_01.jpg',
        title: '【老粗布圆枕】',
        desc: '加大加长 纯棉透气 纯荞麦壳\n 加大加长 纯棉透气 可拆洗',
        navigate: '/pages/items/item01/item01',
      },
      {
        logo: '/images/index_item_02.jpg',
        title: '【老粗布方枕】',
        desc: '加大加长 纯棉透气 纯荞麦壳\n 加大加长 纯棉透气 可拆洗',
        navigate: '/pages/items/item01/item01',
      },
      {
        logo: '/images/index_item_03.jpg',
        title: '【荞麦茶】',
        desc: '荞麦茶是以荞麦为主的茶饮。富含有大量的芸香甙和烟酸（维生素PP）',
        navigate: '/pages/items/item01/item01',
      },

    ]
  },

  onLoad: function () {
  },

  toDetail: function (e) {
    console.log(e);
    var index = e.currentTarget.dataset.index;
    console.log(index);
  },

  // 触发事件：调用拨打电话api
  calling: function () {
    wx.makePhoneCall({
      phoneNumber: app.globalData.phoneNumber,
    })
  },

  onShareAppMessage: function () {
  }


})
