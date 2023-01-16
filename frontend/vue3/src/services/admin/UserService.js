import Api from '@/services/Api'
import secret from '../../Secret'

export default {
    login(data) {
        return Api(`${secret.LARAVEL_APP_URL}`).post(`user/login/`, data)
    },
    profile() {
        return Api(`${secret.LARAVEL_APP_URL}`).get(`user/profile/`,)
    },
    getAllUsers() {
        return Api(`${secret.LARAVEL_APP_URL}`).get(`user/`)
    },
    getUserById(id) {
        return Api(`${secret.LARAVEL_APP_URL}`).get(`user/${id}`)
    },
    createUser(data) {
        return Api(`${secret.LARAVEL_APP_URL}`).post('user/', data)
    },
    updateUser(data, id) {
        return Api(`${secret.LARAVEL_APP_URL}`).post(`user/${id}?_method=PUT`, data)
    },
    changeStatusUser(data, id) {
        return Api(`${secret.LARAVEL_APP_URL}`).put(`user/changeStatus/${id}`, data)
    },
    deleteUserById(id) {
        return Api(`${secret.LARAVEL_APP_URL}`).delete(`user/${id}`)
    }
}
