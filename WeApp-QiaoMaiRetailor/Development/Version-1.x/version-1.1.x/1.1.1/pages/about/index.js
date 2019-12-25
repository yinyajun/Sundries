import {
  $wuxDialog
} from '../../dist/index'

const app = getApp()


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
    phoneNumber: app.globalData.phoneNumber,
    logo: app.globalData.logo,
    list: [{
        id: 'ours',
        name: '关于我们',
        open: false,
        pages: [{
          'eng': 'retailorIntro',
          'ch': '店铺简介',
          'bind': ''
        }],
      },
      {
        id: 'contact',
        name: '联系方式',
        open: false,
        pages: [{
            'eng': 'phoneNumber',
            'ch': '拨打手机',
            'bind': 'onPhoneCall'
          },
          {
            'eng': 'wechatNumber',
            'ch': '添加微信',
            'bind': ''
          },
        ],
      },

      {
        id: 'share',
        name: '分享给朋友',
        open: false,
        pages: [{
            'eng': 'shareApp',
            'ch': '分享小程序',
            'bind': ''
          },
          {
            'eng': 'showCode',
            'ch': '小程序二维码',
            'bind': 'showCode'
          }
        ],
      },
      {
        id: 'feedback',
        name: '服务与支持',
        open: false,
        pages: [{
          'eng': 'feedback',
          'ch': '告诉我们',
          'bind': 'showSupport'
        }]
      }
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

  onLoad() {},
  open() {
    if (this.timeout) clearTimeout(this.timeout)
    this.timeout = setTimeout(hideDialog, 3000)
  },


  showSupport() {
    const alert = (content) => {
      $wuxDialog('#wux-dialog--alert').alert({
        resetOnClose: true,
        title: '提示',
        content: content,
      })
    }

    $wuxDialog().open({
      resetOnClose: true,
      title: '请问需要反馈什么问题？',
      content: '反馈产品问题请点击服务反馈，反馈小程序问题请点技术反馈。',
      verticalButtons: !0,
      buttons: [{
          text: '服务反馈',
          bold: !0,
          onTap(e) {
            alert('目前支持微信联系')
          },
        },
        {
          text: '技术反馈',
          bold: !0,
          onTap(e) {
            alert('目前支持微信联系')
          },
        },
        {
          text: '没啥事',
          bold: !0,
        },
      ],
    })
  },



  kindToggle: function(e) {
    var id = e.currentTarget.id,
      list = this.data.list;
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



  onPhoneCall: function() {
    wx.makePhoneCall({
      phoneNumber: this.data.phoneNumber,
    })
  },

  showCode: function() {
    wx.previewImage({
      urls: [app.globalData.appCodeURL],
    })
  },


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





})