// pages/articles/article01/article01.js
Page({

  /**
   * 页面的初始数据
   */
  data: {

  },

  copyLink:function(e){
    var url=e._relatedInfo.anchorTargetText;
    wx.setClipboardData({
      data: url,
      success: function (res) {
        wx.showToast({
          title: '网址已复制',
          icon: 'success',
          duration: 3000
        });
      },
    })

  },
  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})