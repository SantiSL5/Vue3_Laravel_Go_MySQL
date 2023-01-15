import store from '../../store';
import Constant from '../../Constant';
// import UserService from '../client/UserService';
import UserServiceAdmin from '../client/UserService';

export default {

    async authGuardAdmin(to, from, next) {
        try {
            if (localStorage.getItem('tokenAdmin')) {
                const response = await UserServiceAdmin.profile();
                if (response.status === 200) {
                    next();
                }
            } else {
                next('/login');
            }
        } catch (error) {
            store.dispatch(`user/${Constant.LOGOUT}`);

        }
    },

    async AuthGuard(to, from, next) {
        if (localStorage.getItem('token')) {
            await store.dispatch(`userClient/${Constant.PROFILE_USER}`);
        }
        if (store.getters['userClient/getAuth'] && localStorage.getItem('token')) {
            next();
        } else {
            next('/home');
        }
    },

    noAuthGuard(to, from, next) {
        // if (!store.getters['user/getAuth'] && !localStorage.getItem('isAuth')) {
        if (!localStorage.getItem('token')) {
            next();
        } else {
            next('/home');
        }
    },
}
