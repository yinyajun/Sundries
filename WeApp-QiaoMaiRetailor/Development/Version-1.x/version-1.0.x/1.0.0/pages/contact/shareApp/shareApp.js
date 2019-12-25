// pages/contact/shareApp/shareApp.js
const app= getApp()

Page({

  /**
   * 页面的初始数据
   */
  data: {
    retailName: app.globalData.retailName,
    wechatNumber: app.globalData.wechatNumber,
    phoneNumber: app.globalData.phoneNumber
  },



  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {
    return {
      title: '英英荞麦，常州市金坛区更好的荞麦零售。',
      path: '/pages/index/index',
      imageUrl: "/images/my_avatar.png",
      success: (res) => {
        console.log("转发成功", res);
      },
      fail: (res) => {
        console.log("转发失败", res);
      }
    }
  }
})