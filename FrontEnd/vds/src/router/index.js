import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const Category        = () => import('@/views/Category.vue');
const DFS             = () => import('@/views/graph/DFS.vue');
const DFSOnce         = () => import('@/views/graph/DFSOnce.vue');

const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err);
};

const routes = [
  { path: '/', component: Category },
  { path: '/graph/dfs', component: DFS },
  { path: '/graph/dfs-once', component: DFSOnce },
];

const router = new VueRouter({
  routes
});

export default router;
