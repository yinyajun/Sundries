//app.js
App({
  onLaunch: function() {},
  globalData: {
    retailName: '英英荞麦',
    phoneNumber: '18362292595',
    wechatNumber: 'wuly108',
    appCodeURL: 'cloud://yingyingqiaoma-c46492.7969-yingyingqiaoma-c46492/assets/images/app_code.jpg',

    log:'/assets/images/logo_dark_trans.png',
    personCodeURL:'cloud://yingyingqiaoma-c46492.7969-yingyingqiaoma-c46492/assets/images/person_code.jpg',
    shareFunction: function() {
      return {
        title: '英英荞麦零售小铺位于常州金坛区，提供更加实惠、新鲜、健康的荞麦系列产品，市区内送货上门。',
        path: '/pages/index/index',
        imageUrl: "/assets/images/share_img.jpg",
        success: (res) => {
          console.log("转发成功", res);
        },
        fail: (res) => {
          console.log("转发失败", res);
        }
      }
    },
  },

  onShareAppMessage: function() {
    return {
      title: '英英荞麦零售，金坛区更好的荞麦零售。',
      path: '/pages/index/index',
      imageUrl: "/assets/images/share_img.jpg",
      success: (res) => {
        console.log("转发成功", res);
      },
      fail: (res) => {
        console.log("转发失败", res);
      }
    }
  }




})