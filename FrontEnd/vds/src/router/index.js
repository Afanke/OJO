import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const Category            = () => import('@/components/Category.vue');
const DFS        = () => import('@/components/graph/DFS.vue');
const DFSOnce       = () => import('@/components/graph/DFSOnce.vue');


const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err);
};

const routes = [
  { path: '/', component: Category },
  // { path: '/1234', component: Practice },
  { path: '/graph/dfs', component: DFS },
  { path: '/graph/dfs-once', component: DFSOnce },
  // { path: '/', component:Test }
  // { path: '/', component:Carton }
];

const router = new VueRouter({
  routes
});

export default router;
