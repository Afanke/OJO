import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const Home = () => import('@/views/Home.vue');
const Practice = () => import('@/views/practice/Practice.vue');
const Contest = () => import('@/views/contest/Contest.vue');
const PracticeStatus = () => import('@/views/status/PracticeStatus.vue');
const ContestStatus = () => import('@/views/status/ContestStatus.vue');
const About = () => import('@/views/About.vue');
const PracticeResult = () => import('@/views/practice/Result.vue');
const ContestResult = () => import('@/views/contest/Result.vue');
const Test = () => import('@/views/Test.vue');
const PracticeAnswer = () => import('@/views/practice/Answer.vue');
const ContestAnswer = () => import('@/views/contest/Answer.vue');
const ContestDetail = () => import('@/views/contest/Detail.vue');
const Announcement = () => import('@/views/Announcement.vue');
const PracticeRank = () => import('@/views/rank/PracticeRank.vue');
const UserHome = () => import('@/views/user/Home.vue');
const UserSettings = () => import('@/views/user/Settings.vue');
const ResetPassword = () => import('@/views/user/ResetPassword.vue');
const CaptchaBox = () => import('@/components/CaptchaBox.vue');

const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err);
};

const routes = [
    {path: '/home', component: Announcement},
    {path: '/practice', component: Practice},
    {path: '/practice/answer', component: PracticeAnswer},
    {path: '/contest', component: Contest},
    {path: '/status/practice', component: PracticeStatus},
    {path: '/status/contest', component: ContestStatus},
    {path: '/about', component: About},
    {path: '/practice/result', component: PracticeResult},
    {path: '/contest/result', component: ContestResult},
    {path: '/contest/detail', component: ContestDetail},
    {path: '/contest/answer', component: ContestAnswer},
    {path: '/rank/PracticeRank', component: PracticeRank},
    {path: '/user/home', component: UserHome},
    {path: '/user/settings', component: UserSettings},

    // {path: '/', redirect: '/practice'}
    { path: '/', component:ResetPassword }
    // { path: '/', component:CaptchaBox }
];

const router = new VueRouter({
    routes
});

export default router;
