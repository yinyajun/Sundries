const app = getApp()

const icon = '/assets/images/iconfont-about.png'
const buttons = [{
    label: '联系电话',
    icon: '/assets/images/iconfont-phone-btn.png',
  },
  {
    opentype: 'onClick',
    label: '展示二维码',
    icon: '/assets/images/iconfont-code-btn.png',
  },
  {
    openType: 'share',
    label: '分享店铺',
    icon: '/assets/images/iconfont-share-btn.png',
  },
  {
    openType: 'contact',
    label: '客服会话',
    icon: '/assets/images/iconfont-service-btn.png',
  },
  {
    opentype: 'onClick',
    label: '关于我们',
    icon: '/assets/images/iconfont-about-btn.png',
  },
]

Page({
  //页面初始数据，注意要用绝对路径
  data: {
    imgUrls: [
      '/assets/images/swiper_01.jpg',
      '/assets/images/swiper_02.jpg',
      '/assets/images/swiper_03.jpg',
      '/assets/images/swiper_04.jpg',
    ],
    indicatorDots: true,
    autoplay: true,
    interval: 3000,
    duration: 1000,
    notice: '欢迎光临' + app.globalData.retailName + '零售小铺，我们的产品经济实惠，绿色新鲜，金坛市区内免费送货上门。手机：' + app.globalData.phoneNumber + '，敬请惠存。',
    proList: [{
        logo: '/assets/images/index_item_01.jpg',
        title: '【老粗布圆枕】',
        desc: '加大加长 纯棉透气 纯荞麦壳\n 加大加长 纯棉透气 可拆洗',
        navigate: '/pages/items/item01/item01',
      },
      {
        logo: '/assets/images/index_item_02.jpg',
        title: '【老粗布方枕】',
        desc: '加大加长 纯棉透气 纯荞麦壳\n 加大加长 纯棉透气 可拆洗',
        navigate: '/pages/items/item01/item01',
      },
      {
        logo: '/assets/images/index_item_03.jpg',
        title: '【荞麦茶】',
        desc: '荞麦茶是以荞麦为主的茶饮。富含有大量的芸香甙和烟酸（维生素PP）',
        navigate: '/pages/items/item01/item01',
      },
      {
        logo: '/assets/images/index_item_03.jpg',
        title: '【荞麦粉】',
        desc: '荞麦茶是以荞麦为主的茶饮。富含有大量的芸香甙和烟酸（维生素PP）',
        navigate: '/pages/items/item01/item01',
      },
      {
        logo: '/assets/images/index_item_03.jpg',
        title: '【荞麦米】',
        desc: '荞麦茶是以荞麦为主的茶饮。富含有大量的芸香甙和烟酸（维生素PP）',
        navigate: '/pages/items/item01/item01',
      }
    ],
    types: ['topLeft', 'topRight', 'bottomLeft', 'bottomRight', 'center'],
    typeIndex: 3,
    colors: ['light', 'stable', 'positive', 'calm', 'balanced', 'energized', 'assertive', 'royal', 'dark'],
    colorIndex: 4,
    dirs: ['horizontal', 'vertical', 'circle'],
    dirIndex: 1,
    sAngle: 0,
    eAngle: 360,
    spaceBetween: 10,
    buttons,
  },

  onLoad: function() {},

  onShareAppMessage: app.globalData.shareFunction,



  onClick(e) {
    console.log('onClick', e)
    // button index : 0
    if (e.detail.index === 0) {
      wx.makePhoneCall({
        phoneNumber: app.globalData.phoneNumber,
      })
    }
    // button index : 1
    if (e.detail.index === 1) {
      wx.previewImage({
        urls: [app.globalData.appCodeURL],
      })
    }
    // button index : 4
    if (e.detail.index === 4) {
      wx.switchTab({
        url: '/pages/about/index'
      })
    }
  },
  onContact(e) {
    console.log('onContact', e)
  },
  onGotUserInfo(e) {
    console.log('onGotUserInfo', e)
  },

})