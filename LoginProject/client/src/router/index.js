import Vue from 'vue'
import Router from 'vue-router'
import LoginUserName from '@/components/LoginUserName'
import LoginPassword from '@/components/LoginPassword'
import EmailVerification from '@/components/EmailVerification'
import TwoStepAuth from '@/components/TwoStepAuth'
import RegistrationForm from '@/components/RegistrationForm'
import ForgotPassword from '@/components/ForgotPassword'
import ChangePassword from '@/components/ChangePassword'
import ForgotPasswordOTPVerification from '@/components/ForgotPasswordOTPVerification'
import Home from '@/components/home'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      redirect: '/login'
    },
    {
      path: '/login',
      name: 'LoginUserName',
      component: LoginUserName
    },
    {
      path: '/loginpassword',
      name: 'LoginPassword',
      component: LoginPassword
    },
    {
      path: '/emailverification',
      name: 'EmailVerification',
      component: EmailVerification
    },
    {
      path: '/forgotpasswordotpverification/',
      name: 'ForgotPasswordOTPVerification',
      component: ForgotPasswordOTPVerification
    },
    {
      path: '/twostepverification',
      name: 'TwoStepAuth',
      component: TwoStepAuth
    },
    {
      path: '/register',
      name: 'Register',
      component: RegistrationForm
    },
    {
      path: '/forgotpassword',
      name: 'ForgotPassword',
      component: ForgotPassword
    },
    {
      path: '/changepassword',
      name: 'ChangePassword',
      component: ChangePassword
    },
    {
      path: '/home',
      name: 'Home',
      component: Home
    }
  ]
})
