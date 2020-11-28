import {extenObsevable} from 'mobx';

class UserStore {
    constructor() {
        extenObsevable(this, {

            loading: true,
            isLoggedIn: false,
            username: ''
        })
    }

}

export default new UserStore();