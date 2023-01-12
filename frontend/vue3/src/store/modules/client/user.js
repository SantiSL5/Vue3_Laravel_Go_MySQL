import Constant from "../../../Constant";
import UserService from "../../../services/client/CategoryService";

export const categoryClient = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.GET_ALL_CATEGORIES]: (state, payload) => {
            state.categorieslist = payload;
        },
        [Constant.GET_ONE_CATEGORY]: (state) => {
            state.allUsers;
        },
    },
    actions: {
        [Constant.GET_ALL_CATEGORIES]: (store) => {
            CategoryService.getAllCategories().then(data => {
                store.commit(Constant.GET_ALL_CATEGORIES, data.data);
            });
        },
        [Constant.GET_ONE_CATEGORY]: (store) => {
            CategoryService.getCategoryById().then(data => {
                store.commit(Constant.GET_ONE_CATEGORY, data.data);
            });
        },
    },
    getters: {
        getCategory(state) {
            return state.categorieslist;
        }
    }
}
