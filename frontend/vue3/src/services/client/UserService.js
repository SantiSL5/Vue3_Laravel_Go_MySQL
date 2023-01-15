import Api from '@/services/Api'
import secret from '../../Secret'

export default {
    register(data) {
        return Api(`${secret.GO_APP_URL}`).post('user/register', data)
    },
    login(data) {
        return Api(`${secret.GO_APP_URL}`).post(`user/login`, data)
    },
}