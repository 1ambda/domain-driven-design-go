export default {
    cleanFlashMessage(state: any) {
        state.flashMessage = ''
    },

    setFlashMessage(state: any, message: string) {
        state.flashMessage = message
    },

    logout(state: any) {
        state.uid = ''
    },

    login(state: any, uid: string) {
        state.uid = uid
    },

    changePath(state: any, path: string) {
        state.path = path
    }
}
