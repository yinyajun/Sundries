// pages/articles/article01/article01.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    value: [36],
    dynamic_font:34
  },

  copyLink: function(e) {
    var url = e._relatedInfo.anchorTargetText;
    wx.setClipboardData({
      data: url,
      success: function(res) {
        wx.showToast({
          title: '网址已复制',
          icon: 'success',
          duration: 3000
        });
      },
    })

  },

  onShareAppMessage: function() {},

  onChange(e) {
    console.log('onChange', this.data.dynamic_font)
    this.setData({
      value: e.detail.value,
      dynamic_font: e.detail.value[0] / 20 + 34
      
    })
  },
  afterChange(e) {
    console.log('afterChange', this.data.dynamic_font)
    this.setData({
      value: e.detail.value,
      dynamic_font: e.detail.value[0] / 20 + 34

    })
  },
})