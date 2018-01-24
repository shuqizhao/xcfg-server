/**
 *  路由文件
 *  引入[Vue, VueRouter]模块
 */
import Vue from 'vue'
import VueRouter from 'vue-router'


const XcfgDevList = r => require.ensure([], () => r(require('../biz/xcfg/xcfgDevList.vue')), 'xcfg')
const XcfgTestingList = r => require.ensure([], () => r(require('../biz/xcfg/xcfgTestingList.vue')), 'xcfg')
const XcfgLabsList = r => require.ensure([], () => r(require('../biz/xcfg/xcfgLabsList.vue')), 'xcfg')
const XcfgProdList = r => require.ensure([], () => r(require('../biz/xcfg/xcfgProdList.vue')), 'xcfg')
const XcfgAdd = r => require.ensure([], () => r(require('../biz/xcfg/xcfgAdd.vue')), 'xcfg')
const XcfgApolloAdd = r => require.ensure([], () => r(require('../biz/xcfg/xcfgApolloAdd.vue')), 'xcfg')
const XcfgView = r => require.ensure([], () => r(require('../biz/xcfg/xcfgView.vue')), 'xcfg')
const XcfgHistoryView = r => require.ensure([], () => r(require('../biz/xcfg/xcfgHistoryView.vue')), 'xcfg')


const UserList = r => require.ensure([], () => r(require('../biz/auth/user.vue')), 'auth')
const UserView = r => require.ensure([], () => r(require('../biz/auth/userView.vue')), 'auth')
const UserAdd = r => require.ensure([], () => r(require('../biz/auth/userAdd.vue')), 'auth')
const UserMod = r => require.ensure([], () => r(require('../biz/auth/userMod.vue')), 'auth')

const RoleList = r => require.ensure([], () => r(require('../biz/auth/role.vue')), 'auth')
const RoleView = r => require.ensure([], () => r(require('../biz/auth/roleView.vue')), 'auth')
const RoleAdd = r => require.ensure([], () => r(require('../biz/auth/roleAdd.vue')), 'auth')

const ResourceList = r => require.ensure([], () => r(require('../biz/auth/resource.vue')), 'auth')

const Portal = r => require.ensure([], () => r(require('../biz/portal/clock.vue')), 'portal')


/**
 *  加载模块
 */
Vue.use(VueRouter)

/**
 *  配置路由
 */

const router = new VueRouter({
  routes: [{
    path: '/',
    component: Portal
  },{
    path: '/xcfg/dev',
    component: XcfgDevList
  },{
    path: '/xcfg/testing',
    component: XcfgTestingList
  },{
    path: '/xcfg/labs',
    component: XcfgLabsList
  },{
    path: '/xcfg/prod',
    component: XcfgProdList
  }, {
    path: '/xcfg/xcfgview',
    component: XcfgView
  }, {
    path: '/xcfg/xcfgadd',
    component: XcfgAdd
  }, {
    path: '/xcfg/xcfgapolloadd',
    component: XcfgApolloAdd
  }, {
    path: '/xcfg/xcfgHistoryView',
    component: XcfgHistoryView
  }, {
    path: '/auth/user',
    component: UserList
  }, {
    path: '/auth/userview',
    component: UserView
  }, {
    path: '/auth/useradd',
    component: UserAdd
  } ,{
    path: '/auth/usermod',
    component: UserMod
  }, {
    path: '/auth/role',
    component: RoleList
  }, {
    path: '/auth/roleview',
    component: RoleView
  }, {
    path: '/auth/roleadd',
    component: RoleAdd
  }, {
    path: '/auth/resource',
    component: ResourceList
  }]
})

export default router
