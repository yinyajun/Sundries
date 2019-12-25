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
        logo: 'cloud://yingyingqiaoma-c46492.7969-yingyingqiaoma-c46492/assets/images/index_item_01.jpg',
        title: '【纯棉老粗布圆枕】',
        desc: '荞麦壳枕芯 可拆洗 \n 颈椎牵引 放松斜方肌',
        navigate: '/pages/items/item01/item01',
      },
      {
        logo: 'cloud://yingyingqiaoma-c46492.7969-yingyingqiaoma-c46492/assets/images/index_item_05.jpg',
        title: '【纯棉老粗布信封式大枕】',
        desc: '荞麦壳枕芯 可拆洗 \n多种尺寸可选',
        navigate: '/pages/items/item01/item01',
      },
      {
        logo: 'cloud://yingyingqiaoma-c46492.7969-yingyingqiaoma-c46492/assets/images/index_item_07.jpg',
        title: '【纯棉老粗布花边儿童枕】',
        desc: '荞麦壳枕芯 可拆洗 \n精致花边 儿童专用尺寸',
        navigate: '/pages/items/item01/item01',
      },
      {
        logo: 'cloud://yingyingqiaoma-c46492.7969-yingyingqiaoma-c46492/assets/images/index_item_02.jpg',
        title: '【纯棉老粗布方枕】',
        desc: '荞麦壳枕芯 可拆洗 ',
        navigate: '/pages/items/item01/item01',
      },
      {
        logo: 'cloud://yingyingqiaoma-c46492.7969-yingyingqiaoma-c46492/assets/images/index_item_03.jpg',
        title: '【优质苦荞茶】',
        desc: '纯荞麦香 清香醇厚\n 健胃消食 功效良多',
        navigate: '/pages/items/item01/item01',
      },
      {
        logo: 'cloud://yingyingqiaoma-c46492.7969-yingyingqiaoma-c46492/assets/images/index_item_04.jpg',
        title: '【优质荞麦粉】',
        desc: '无添加剂 保证新鲜\n口感好 可烹制多种美食',
        navigate: '/pages/items/item01/item01',
      },
      {
        logo: 'cloud://yingyingqiaoma-c46492.7969-yingyingqiaoma-c46492/assets/images/index_item_06.jpg',
        title: '【生黑苦荞麦米】',
        desc: '营养价值高 气味清香',
        navigate: '/pages/items/item01/item01',
      },

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