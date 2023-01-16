import Constant from "../../../Constant";
import UserService from "../../../services/client/UserService";
import UserServiceAdmin from "../../../services/admin/UserService";
import router from '../../../router/index';
import { createToaster } from "@meforma/vue-toaster"; // Debería estar en el componente

export const userClient = {
    namespaced: true,

    state: {
        user: {
            username: "",
            email: "",
            img: "",
            type: "",
        },
    },
    mutations: {
        [Constant.REGISTER]: (state, payload) => {
            const toaster = createToaster({ position: "top" });
            toaster.success("User " + payload.username + " successfully registered");
        },
        [Constant.LOGIN]: (state, payload) => {
            localStorage.setItem("token", payload.token);
            state.user = {
                username: payload.username,
                email: payload.email,
                image: payload.image,
                type: 'client',
            };
            router.push({ name: 'home' });
        },
        [Constant.LOGIN_ADMIN]: (state, payload) => {
            localStorage.setItem("tokenAdmin", payload.token);
            state.user = {
                username: payload.user.username,
                email: payload.user.email,
                image: payload.user.image,
                type: 'admin',
            };
            router.push({ name: 'home' });
        },
        [Constant.LOGOUT]: (state) => {
            state.user = {
                username: "",
                email: "",
                image: "",
                type: "",
            };
            router.push({ name: 'home' });
        },
        [Constant.PROFILE_USER]: (state, payload) => {
            state.user = {
                username: payload.username,
                email: payload.email,
                image: payload.image,
                type: payload.type,
            };
            // if (!payload.redirect) {
            //     router.push({ name: 'home' });
            // }
        },
    },
    actions: {
        [Constant.REGISTER]: (store, payload) => {
            UserService.register(payload).then(data => {
                if (data.status == 200) {
                    store.commit(Constant.REGISTER, payload);
                }
            })
            // .catch(e => {
            //     const toaster = createToaster({
            //         position: "top"
            //     });
            //     console.log(e);
            //     if (e["response"]["data"] == "Email is registered") {
            //         toaster.error("Email already exists");
            //     } else {
            //         toaster.error("Error register");
            //     }
            // });
        },
        [Constant.LOGIN]: (store, payload) => {
            const toaster = createToaster({ position: "top" });
            UserService.login(payload).then(data => {
                if (data.data.type == "admin") {
                    UserServiceAdmin.login(payload).then(data => {
                        toaster.success("Admin " + data.data.user.username + " loged successfully");
                        store.commit(Constant.LOGIN_ADMIN, data.data);
                    }).catch(function (error) {
                        console.log(error);
                        toaster.error("Error login Admin");
                    });
                } else {
                    store.commit(Constant.LOGIN, data.data);
                    toaster.success(data.data.username + " loged successfully");
                }
            }).catch(function (error) {
                if (error.response.data == "Incorrect password") {
                    toaster.error(error.response.data);
                } else {
                    toaster.error(error.response.data);
                }
            });

        },
        [Constant.PROFILE_USER]: (store) => {
            if (localStorage.getItem('tokenAdmin')) {
                UserServiceAdmin.profile().then(data => {
                    store.commit(Constant.PROFILE_USER, data.data);
                }).catch(function () {});    
            }else{
                UserService.profile().then(data => {
                    store.commit(Constant.PROFILE_USER, data.data);
                }).catch(function () {});    
            }
        },
        [Constant.LOGOUT]: (store) => {
            localStorage.removeItem("token");
            localStorage.removeItem("tokenAdmin");
            store.commit(Constant.LOGOUT);
        },

    },
    getters: {
        getUser(state) {
            return state.userslist;
        },
        getAuth(state) {
            console.log(state.user);
            return state.user;
        },
    }
}
