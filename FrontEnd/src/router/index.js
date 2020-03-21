import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const Home            = () => import('@/components/Home.vue');
const Practice        = () => import('@/components/practice/Practice.vue');
const Contest         = () => import('@/components/contest/Contest.vue');
const PracticeStatus  = () => import('@/components/practice/Status.vue');
const Rank            = () => import('@/components/Rank.vue');
const About           = () => import('@/components/About.vue');
const PracticeResult  = () => import('@/components/practice/Result.vue');
const ContestResult   = () => import('@/components/contest/Result.vue');
const Test            = () => import('@/components/Test.vue');
const PracticeAnswer  = () => import('@/components/practice/Answer.vue');
const ContestAnswer   = () => import('@/components/contest/Answer.vue');
const ContestDetail   = () => import('@/components/contest/Detail.vue');
const Carton   = () => import('@/components/Carton.vue');

const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err);
};

const routes = [
  { path: '/home', component: Home },
  { path: '/practice', component: Practice },
  { path: '/practice/answer', component: PracticeAnswer},
  { path: '/contest', component: Contest },
  { path: '/status', component: PracticeStatus },
  { path: '/rank', component: Rank },
  { path: '/about', component: About },
  { path: '/practice/result', component: PracticeResult },
  { path: '/contest/result', component: ContestResult },
  { path: '/test', component: Test },
  { path: '/contest/detail', component: ContestDetail },
  { path: '/contest/answer', component: ContestAnswer },

  // { path: '/', redirect: '/contest' }
  // { path: '/', component:Test }
  { path: '/', component:Carton }
];
// routes.afterEach(function (to) {
//   document.title = to.name
//   })
const router = new VueRouter({
  routes
});

// router.beforeEach((to, from, next) => {
//   if (to.path === '/login') return next()
//   const tokenStr = window.sessionStorage.getItem('token')
//   if (!tokenStr) return next('/login')
//   next()
// })
export default router;
