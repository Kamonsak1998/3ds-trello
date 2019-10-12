export const user = {
    namespaced: true,
    actions: {
        getIduser: ({ commit }, idMember) => {
            commit('addIduser', idMember)
        },
        getUsername: ({ commit }, Username) => {
            commit('addUsername', Username)
        },
        getNameBoard: ({ commit }, nameBoard) => {
            commit('addNameBoard', nameBoard)
        },
        getBoard: ({ commit }, idBoard) => {
            commit('addBoard', idBoard)
        },
        getToken: ({ commit }, token) => {
            commit('addToken', token)
        },
        getAvatar:({commit},urlAvatar) =>{
            commit('setAvatar',urlAvatar)
        },
        logout: ({ commit }) => {
            commit('remove')
        }
    }, state: {
        idMember: null,
        username: null,
        nameBoard: null,
        idBoard: null,
        urlAvatar: null,
        token: null,
    }, mutations: {
        addIduser: (state, idMember) => {
            state.idMember = idMember
        },
        addUsername: (state, newUsername) => {
            state.username = newUsername
        },
        addNameBoard: (state, newNameBoard) => {
            state.nameBoard = newNameBoard
        },
        addBoard: (state, idBoard) => {
            state.idBoard = idBoard
        },
        addToken: (state, newToken) => {
            state.token = newToken
            localStorage.setItem('token', newToken)
        },
        setAvatar:(state, urlAvatar) => {
            state.urlAvatar = urlAvatar
        },
        remove: (state) => {
            state.idBoard = null
            state.nameBoard = null
            state.username = null
            state.idMember = null
            state.token = null
            state.urlAvatar = null
            localStorage.removeItem('token', null)
        }
    }, getters: {
        idMember: state => state.idMember,
        username: state => state.username,
        nameBoard: state => state.nameBoard,
        idBoard: state => state.idBoard,
        token: state => state.token,
        urlAvatar: state => state.urlAvatar
    }
}