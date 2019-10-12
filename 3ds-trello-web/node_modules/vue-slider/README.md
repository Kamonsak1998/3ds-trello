# vue-slider
Vue 的图片轮播组件 (a vue slider component)，可以直接使用，支持配置

本项目是 [使用 React 实现一个轮播组件](http://qiutc.me/post/%E4%BD%BF%E7%94%A8-React-%E5%AE%9E%E7%8E%B0%E4%B8%80%E4%B8%AA%E8%BD%AE%E6%92%AD%E7%BB%84%E4%BB%B6.html) 的 Vue 版本。
# 使用方法
- 用 npm 安装本组件

```
npm i vue-slider
```
- 在项目中引用配置示例

```
<template lang="jade">
  div
    slider(:items="items", :speed="2", :delay="3", :pause="true", :autoplay:"true", :dots="true", :arrows:"true")
</template>
<script>
  import Slider from '../slider/Slider'
  export default {
    data () {
      return {
        items: [
          {
            src: 'http://placehold.it/350x150/999999/000000',
            alt: 'images-1'
          },
          {
            src: 'http://placehold.it/350x150/eeeeee/c90023',
            alt: 'images-2'
          },
          {
            src: 'http://placehold.it/350x150/333333/eeeeee',
            alt: 'images-3'
          }
        ]
      }
    }, //图片数组，有几张图片放几张
    components: {
      Slider
    }
  }
</script>
```

# 组件参数

以下参数除了 `items` 必须传入，其他都有默认值，可以查看 `slider/Slider.vue` 文件中的 `props` 设置

参数 | 是否必须 | 参数值 | 默认值
----|------|----|----
items | 是 | 图片数组 | 无
speed  | 否 |   轮播切换图片的速度 | 1.5
delay  | 否 |  自动轮播时候停留的时间 | 2
pause  | 否 |  鼠标放上图片是否停止自动轮播 | true
autoplay | 否 |  是否自动轮播 | true    
dots  | 否 |  是否需要下面的轮播点点位 | true    
arrows  | 否 |  是否需要左右点击箭头 | true  

# example 中的示例
安利一个快速配置 webpack 的工具：[:hammer:tooling](https://github.com/egoist/tooling)
```
// 安装 webpack 工具 tooling
npm i -g tooling

// 安装本组件依赖
npm i -d

// 查看例子
npm run watch:example

// 打包例子
npm run build:example

// 打包 slider 组件
npm run build

```
在浏览器打开 http://localhost:3003/。

打包的示例文件在 `builid` 下面。

打包的 `slider` 组件 `js` 在 `dist` 下面。