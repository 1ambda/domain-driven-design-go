import Vue from 'vue'
import Vuex from 'vuex'

import mutations from '@/store/mutation.ts'
import actions from '@/store/action.ts'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        uid: '', // indicates user is logged in
        flashMessage: '',
        path: '',
    },
    getters: {
        authenticated: (state: any) => {
            return state.uid !== ''
        },
        uid: (state: any) => {
            return state.uid
        },
    },
    mutations,
    actions,
})

