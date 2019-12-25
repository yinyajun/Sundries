// pages/contact/phoneNumber/phoneNumber.js
const app = getApp()

Page({

  /**
   * 页面的初始数据
   */
  data: {
    phoneNumber: app.globalData.phoneNumber
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    wx.makePhoneCall({phoneNumber: this.data.phoneNumber})
    wx.navigateBack({delta: -1})
  },


})