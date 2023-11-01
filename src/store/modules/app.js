import Cookies from 'js-cookie'
import {listType} from "@/api/app";

const state = {
  sidebar: {
    opened: Cookies.get('sidebarStatus') ? !!+Cookies.get('sidebarStatus') : true,
    withoutAnimation: false
  },
  device: 'desktop',
  form: {
    prdEnv: false,
    schUserName: ''
  }
}

const mutations = {
  TOGGLE_SIDEBAR: state => {
    state.sidebar.opened = !state.sidebar.opened
    state.sidebar.withoutAnimation = false
    if (state.sidebar.opened) {
      Cookies.set('sidebarStatus', 1)
    } else {
      Cookies.set('sidebarStatus', 0)
    }
  },
  CLOSE_SIDEBAR: (state, withoutAnimation) => {
    Cookies.set('sidebarStatus', 0)
    state.sidebar.opened = false
    state.sidebar.withoutAnimation = withoutAnimation
  },
  TOGGLE_DEVICE: (state, device) => {
    state.device = device
  },
  CHANGE_ENV: (state, env) => {
    state.form.prdEnv = env
  },
  CHANGE_SUN: (state, suname) => {
    console.log('app.')
    console.log(state.form.schUserName)
    state.form.schUserName = suname
  }
}

const actions = {
  toggleSideBar({ commit }) {
    commit('TOGGLE_SIDEBAR')
  },
  closeSideBar({ commit }, { withoutAnimation }) {
    commit('CLOSE_SIDEBAR', withoutAnimation)
  },
  toggleDevice({ commit }, device) {
    commit('TOGGLE_DEVICE', device)
  },
  changeEnv({ commit }, env) {
    commit('CHANGE_ENV', env)
  },
  changeSchUserName({ commit }, suname) {
    commit('CHANGE_SUN', suname)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
