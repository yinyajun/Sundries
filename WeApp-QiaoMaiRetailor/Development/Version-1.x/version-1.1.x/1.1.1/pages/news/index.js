import {
  tab1_newsList
} from './tab1.js'
import {
  tab2_newsList
} from './tab2.js'


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
    label: '回到首页',
    icon: '/assets/images/iconfont-index-btn.png',
  },
]

Page({
  data: {
    retailName: app.globalData.retailName,
    current: 'tab1',
    tab1_newsList,
    tab2_newsList,

    tabs: [{
        key: 'tab1',
        title: '店铺动态',
      },
      {
        key: 'tab2',
        title: '小知识',
      },
      {
        key: 'tab3',
        title: '活动推广',
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


  onChange(e) {
    console.log('onChange', e)
    this.setData({
      current: e.detail.key,
    })
  },


  onTabsChange(e) {
    console.log('onTabsChange', e)
    const {
      key
    } = e.detail
    const index = this.data.tabs.map((n) => n.key).indexOf(key)

    this.setData({
      key,
      index,
    })
  },


  onSwiperChange(e) {
    console.log('onSwiperChange', this.data.index, e)
    const {
      current: index,
      source
    } = e.detail
    const {
      key
    } = this.data.tabs[index]

    if (!!source) {
      this.setData({
        key,
        index,
      })
    }
  },

  loadMore1: function(e) {
    console.log('loadMore1', e);
    this.getData();
  },

  loadMore2: function(e) {
    console.log('loadMore2', e);
    this.getData();
  },

  onReachBottom() {
    console.log('触底了');

  },


  getData() {
    wx.showLoading({
      title: "更新数据"
    });
    setTimeout(
      () => {
        // for (var i = 0; i < 10; i++) {
        //   listData.push(_item)
        // }
        // this.setData({ listData, loadTips: '上拉加载更多' });
        // this.autoHeight()
        wx.hideLoading();
      }, 1000)
  },

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
        url: '/pages/index/index'
      })
    }
  },



  onContact(e) {
    console.log('onContact', e)
  },
  onGotUserInfo(e) {
    console.log('onGotUserInfo', e)
  },

  onShareAppMessage: app.globalData.shareFunction,

  showArticle01: function() {
    wx.navigateTo({
      url: '/pages/articles/article01/article01',
    })
  },




  showArticle03: function() {
    wx.navigateTo({
      url: '/pages/articles/article03/article03',
    })

  },





})