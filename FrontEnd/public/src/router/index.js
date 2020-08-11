import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const Home            = () => import('@/views/Home.vue');
const Practice        = () => import('@/views/practice/Practice.vue');
const Contest         = () => import('@/views/contest/Contest.vue');
const PracticeStatus  = () => import('@/views/practice/Status.vue');
const Rank            = () => import('@/views/Rank.vue');
const About           = () => import('@/views/About.vue');
const PracticeResult  = () => import('@/views/practice/Result.vue');
const ContestResult   = () => import('@/views/contest/Result.vue');
const Test            = () => import('@/views/Test.vue');
const PracticeAnswer  = () => import('@/views/practice/Answer.vue');
const ContestAnswer   = () => import('@/views/contest/Answer.vue');
const ContestDetail   = () => import('@/views/contest/Detail.vue');
const Announcement    = () => import('@/views/Announcement.vue');
const ACMRank    = () => import('@/views/rank/ACMRank.vue');
const Carton   = () => import('@/views/Carton.vue');
const UserHome   = () => import('@/views/user/Home.vue');
const UserSettings   = () => import('@/views/user/Settings.vue');

const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err);
};

const routes = [
  { path: '/home', component: Announcement },
  { path: '/practice', component: Practice },
  { path: '/practice/answer', component: PracticeAnswer},
  { path: '/contest', component: Contest },
  { path: '/status', component: PracticeStatus },
  { path: '/rank', component: Rank },
  { path: '/about', component: About },
  { path: '/practice/result', component: PracticeResult },
  { path: '/contest/result', component: ContestResult },
  { path: '/contest/detail', component: ContestDetail },
  { path: '/contest/answer', component: ContestAnswer },
  { path: '/rank/ACMRank', component: ACMRank },
  // { path: '/test', component: Test },
  { path: '/user/home', component: UserHome },
  { path: '/user/settings', component: UserSettings },

  { path: '/', redirect: '/practice' }
  // { path: '/', component: UserSettings }
  // { path: '/', component:Test }
];
// routes.afterEach(function (to) {
//   document.title = to.name
//   })
const router = new VueRouter({
  routes
});

// router.beforeEach((to, from, next) => {
  
  
// })
export default router;
