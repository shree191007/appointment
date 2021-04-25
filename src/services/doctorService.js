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

    
   
}}
export default doctorServices