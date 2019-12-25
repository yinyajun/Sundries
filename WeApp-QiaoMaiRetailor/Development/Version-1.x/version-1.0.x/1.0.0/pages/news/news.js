// pages/news/news.js
const app = getApp()

Page({

  /**
   * 页面的初始数据
   */
  data: {
    retailName: app.globalData.retailName,
    articleList: [
      {
        title: '荞麦枕和健康1',
        desc: '加大加长 纯棉透气 纯荞麦壳\n 加大加长 纯棉透气 可拆洗',
        navigate: "pages/articles/article01/article01",
      },
      {
        title: '荞麦枕和健康2',
        desc: '加大加长 纯棉透气 纯荞麦壳\n 加大加长 纯棉透气 可拆洗\n 加大加长 纯棉透气 可拆洗',
        navigate: "pages/articles/article02/article02",
      },
      {
        title: '荞麦枕和健康3',
        desc: '荞麦茶是以荞麦为主的茶饮。富含有大量的芸香甙和烟酸（维生素PP）',
        navigate: "pages/articles/article03/article03",
      },
    ],
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {

  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})