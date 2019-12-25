// pages/contact/contact.js

const app = getApp()

Page({

  /**
   * 页面的初始数据
   */
  data: {
    title: app.globalData.retailName + '(' + app.globalData.phoneNumber + ')',
    desc: '常州金坛 | 送货上门 \n 荞麦系列 | 经济实惠 \n 手工制作 | 绿色新鲜 \n',

    list: [
      {
        id: 'feedback',
        name: '关于' + app.globalData.retailName,
        open: false,
        pages: [{ 'eng': 'retailorIntro', 'ch': '店铺简介' }],
      },
      {
        id: 'widget',
        name: '联系方式',
        open: false,
        pages: [
          { 'eng': 'phoneNumber', 'ch': '拨打手机' },
          { 'eng': 'wechatNumber', 'ch': '添加微信' },
        ],
      },

      {
        id: 'nav',
        name: '分享给朋友',
        open: false,
        pages: [
          { 'eng': 'shareApp', 'ch': '分享小程序' },
          { 'eng': 'showCode', 'ch':'获得小程序二维码' }],
      },
      {
        id: 'search',
        name: '技术支持',
        open: false,
        pages: [
          { 'eng': 'feedback', 'ch': '小程序技术反馈' }]
      }
    ],

  },

  kindToggle: function (e) {
    var id = e.currentTarget.id, list = this.data.list;
    for (var i = 0, len = list.length; i < len; ++i) {
      if (list[i].id == id) {
        list[i].open = !list[i].open
      } else {
        list[i].open = false
      }
    }
    this.setData({
      list: list
    });
  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})