import Api from '@/services/Api'
import secret from '../../Secret'

export default {
    getAllCategories() {
        return Api(`${secret.GO_APP_URL}`).get(`category/`)
    },
    getCategoryById(id) {
        return Api(`${secret.GO_APP_URL}`).get(`category/${id}`)
    },
    createCategory(data) {
        return Api(`${secret.GO_APP_URL}`).post('category/', data)
    },
    updateCategory(data, id) {
        return Api(`${secret.GO_APP_URL}`).post(`category/${id}?_method=PUT`, data)
    },
    changeStatusCategory(data, id) {
        return Api(`${secret.GO_APP_URL}`).put(`category/changeStatus/${id}`, data)
    },
    deleteCategoryById(id) {
        return Api(`${secret.GO_APP_URL}`).delete(`category/${id}`)
    }
}
