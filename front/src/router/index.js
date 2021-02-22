import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Setting from '@/components/Setting'
import File from '@/components/File'
import List from "@/components/List";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'list',
      component: List
    },
    {
      path: '/setting',
      name: 'setting',
      component: Setting
    },
    {
      path: '/file',
      name: 'file',
      component: File
    }
  ]
})
