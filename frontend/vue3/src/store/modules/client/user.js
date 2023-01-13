import Constant from "../../../Constant";
import UserService from "../../../services/client/UserService";
// import { createToaster } from "@meforma/vue-toaster"; // Debería estar en el componente

export const userClient = {
    namespaced: true,

    state: {
        user: {
            is_active: false,
            username: "",
            email: "",
            img: "",
            type: "",
        },
    },
    mutations: {
        [Constant.REGISTER]: (state, payload) => {
            state.categorieslist = payload;
        },
        [Constant.LOGIN]: (state) => {
            state.allUsers;
        },
        [Constant.LOGOUT]: (state) => {
            state.user = {
                is_active: "",
                username: "",
                email: "",
                img: "",
                type: "",
            };
            // router.push({ name: 'home' });
        },
    },
    actions: {
        [Constant.REGISTER]: (store, payload) => {
            console.log(payload);
            UserService.register(payload).then(data => {
                console.log(data);
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
        [Constant.LOGIN]: (store) => {
            UserService.login().then(data => {
                store.commit(Constant.GET_ONE_CATEGORY, data.data);
            });
        },
        [Constant.LOGOUT]: (store) => {
            localStorage.removeItem("token");
            localStorage.removeItem("token_admin");
            store.commit(Constant.LOGOUT);
        },

    },
    getters: {
        getUser(state) {
            return state.userslist;
        }
    }
}
