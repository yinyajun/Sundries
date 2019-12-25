// pages/contact/wechatNumber/wechatNumber.js
const app = getApp()

Page({

  /**
   * 页面的初始数据
   */
  data: {
    retailName: app.globalData.retailName,
    wechatNumber: app.globalData.wechatNumber,
    phoneNumber: app.globalData.phoneNumber
  },


  copyWechatNumber: function() {
    wx.setClipboardData({
      data: this.data.wechatNumber,
      success: function(res) {
        wx.showToast({
          title: '微信号已复制',
          icon: 'success',
          duration: 3000
        });
      },
      // fail: function(res) {},
      // complete: function(res) {},
    })
  },

  copyPhoneNumber: function () {
    wx.setClipboardData({
      data: this.data.phoneNumber,
      success: function (res) {
        wx.showToast({
          title: '手机号已复制',
          icon: 'success',
          duration: 3000
        });
      },
    })
  },

  saveCode: function () {

  },

  

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})