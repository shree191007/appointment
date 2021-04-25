import axios from 'axios'
import config from './../config'
var controller ='/doctors'
var baseurl = config.API_SERVER+controller


var doctorServices={
    getAlldoctors(){
        console.log(baseurl)
        return axios.get(baseurl)
        .then(
            response => response.data)
        .catch(err => console.log(err))

<<<<<<< HEAD
    
   
}}
=======
    }
}
>>>>>>> 265761892482495c332c5ffe91e113e7e6e4615b
export default doctorServices