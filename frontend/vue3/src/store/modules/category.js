import Constant from "../../Constant";
import CategoryService from "../../services/CategoryService";

export const category = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.GET_ALL_CATEGORIES]: (state) => {
            state.allUsers;
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
            return state.category;
        }
    }
}
