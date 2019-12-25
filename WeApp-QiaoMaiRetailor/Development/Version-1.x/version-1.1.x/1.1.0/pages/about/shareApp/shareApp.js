// pages/contact/shareApp/shareApp.js
const app = getApp()

import Poster from '../../../miniprogram_dist/poster/poster';

const posterConfig = {
  jdConfig: {
    width: 750,
    height: 1334,
    backgroundColor: '#fff',
    debug: false,
    blocks: [{
        width: 690,
        height: 808,
        x: 30,
        y: 183,
        borderWidth: 2,
        borderColor: '#f0c2a0',
        borderRadius: 20,
      },
      {
        width: 634,
        height: 74,
        x: 59,
        y: 770,
        backgroundColor: '#fff',
        opacity: 0.5,
        zIndex: 100,
      },
    ],
    texts: [{
        x: 113,
        y: 61,
        baseLine: 'middle',
        text: '英英',
        fontSize: 32,
        color: '#8d8d8d',
      },
      {
        x: 30,
        y: 113,
        baseLine: 'top',
        text: '金坛的小伙伴们，向您推荐',
        fontSize: 38,
        color: '#080808',
      },
      {
        x: 92,
        y: 810,
        fontSize: 38,
        baseLine: 'middle',
        text: '英英荞麦零售',
        width: 570,
        lineNum: 1,
        color: '#8d8d8d',
        zIndex: 200,
      },
      {
        x: 59,
        y: 895,
        baseLine: 'middle',
        text: [{
            text: '英英荞麦零售',
            fontSize: 28,
            color: '#ec1731',
          },
          {
            text: '联系电话：18362292595',
            fontSize: 36,
            color: '#ec1731',
            marginLeft: 30,
          }
        ]
      },
      {
        x: 59,
        y: 945,
        baseLine: 'middle',
        text: [{
            text: '经济实惠',
            fontSize: 28,
            color: '#929292',
          },
          {
            text: '绿色新鲜',
            fontSize: 28,
            color: '#929292',
            marginLeft: 50,
          },
          {
            text: '市区送货上门',
            fontSize: 28,
            color: '#929292',
            marginLeft: 50,
          },
        ]
      },
      {
        x: 360,
        y: 1065,
        baseLine: 'top',
        text: '长按识别小程序码',
        fontSize: 38,
        color: '#080808',
      },
      {
        x: 360,
        y: 1123,
        baseLine: 'top',
        text: '金坛区更好的荞麦零售',
        fontSize: 28,
        color: '#929292',
      },
    ],

    

    images: [{
        // 头像
        width: 62,
        height: 62,
        x: 30,
        y: 30,
        borderRadius: 62,
        url: 'https://lc-I0j7ktVK.cn-n1.lcfile.com/02bb99132352b5b5dcea.jpg',
      },
      {
        //主题图片
        width: 634,
        height: 634,
        x: 59,
        y: 210,
        url: 'https://lc-I0j7ktVK.cn-n1.lcfile.com/193256f45999757701f2.jpeg',
      },
      {
        //二维码
        width: 220,
        height: 220,
        x: 92,
        y: 1020,
        url: 'https://lc-I0j7ktVK.cn-n1.lcfile.com/d719fdb289c955627735.jpg',
      },
      {
        //footer
        width: 750,
        height: 90,
        x: 0,
        y: 1244,
        url: 'https://lc-I0j7ktVK.cn-n1.lcfile.com/67b0a8ad316b44841c69.png',
      }
    ]

  }
}


Page({
  data: {
    posterConfig: posterConfig.jdConfig,
    retailName: app.globalData.retailName,
    wechatNumber: app.globalData.wechatNumber,
    phoneNumber: app.globalData.phoneNumber
  },


  onPosterSuccess(e) {
    const {
      detail
    } = e;
    wx.previewImage({
      current: detail,
      urls: [detail]
    })
  },
  onPosterFail(err) {
    console.error(err);
  },

  /**
   * 异步生成海报
   */
  onCreatePoster() {
    this.setData({
      posterConfig: posterConfig.demoConfig
    }, () => {
      Poster.create(true); // 入参：true为抹掉重新生成
    });
  },

  onCreateOtherPoster() {
    this.setData({
      posterConfig: posterConfig.jdConfig
    }, () => {
      Poster.create(true); // 入参：true为抹掉重新生成 
    });
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


