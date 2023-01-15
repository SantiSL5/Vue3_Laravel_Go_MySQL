import Constant from "../../../Constant";
import UserService from "../../../services/admin/UserService";
import { createToaster } from "@meforma/vue-toaster"; // Debería estar en el componente

export const userAdmin = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.CREATE_ONE_USER]: (state, payload) => {
            state.userslist.push({ ...payload });
        },
        [Constant.GET_ALL_USERS]: (state, payload) => {
            state.userslist = payload;
        },
        [Constant.UPDATE_ONE_USER]: (state, payload) => {
            let index = state.userslist.findIndex(
                item => item.id == payload.id
            );
            state.userslist[index] = payload;
        },
        [Constant.DELETE_ONE_USER]: (state, payload) => {
            let index = state.userslist.findIndex(
                (item) => item.id === payload
            );
            state.userslist.splice(index, 1);
        },
    },
    actions: {
        [Constant.CREATE_ONE_USER]: (store, payload) => {
            UserService.createUser(payload.user).then(data => {
                const toaster = createToaster({
                    position: "top"
                });
                if (data.statusText == "Created") {
                    toaster.success(`User created successfully`);
                    store.commit(Constant.CREATE_ONE_USER, data.data);
                } else {
                    toaster.error(data.data.Message);
                }
            });
        },
        [Constant.UPDATE_ONE_USER]: (store, payload) => {
            UserService.updateUser(payload.user, payload.id).then(data => {
                const toaster = createToaster({
                    position: "top"
                });
                if (data.data.id) {
                    toaster.success(`User updated successfully`);
                    store.commit(Constant.UPDATE_ONE_USER, data.data)
                } else {
                    toaster.error(`User already exists`);
                }
            });
        },
        [Constant.GET_ALL_USERS]: (store) => {
            UserService.getAllUsers().then(data => {
                store.commit(Constant.GET_ALL_USERS, data.data);
            });
        },
        [Constant.GET_ONE_USER]: (store) => {
            UserService.getUserById().then(data => {
                store.commit(Constant.GET_ONE_USER, data.data);
            });
        },
        [Constant.DELETE_ONE_USER]: (store, payload) => {
            UserService.deleteUserById(payload.id).then(
                store.commit(Constant.DELETE_ONE_USER, payload.id)
            );
        },
    },
    getters: {
        getUser(state) {
            return state.userslist;
        }
    }
}
