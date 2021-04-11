<template>
<v-app>
  <v-data-table
    :headers="Headers"
    :items="doctors"
    :items-per-page="5"
    class="elevation-1"
  ></v-data-table>
</v-app>
</template>
<script>
import doctorServices from "@/services/doctorService"

  export default {
    data :() => ({
        Headers:[
            {text:'doctor id', value:"id"},
            {text:'doctor name', value:"name"},
            {text:'type of doctor', value:"type_of_doctor"},
            {text:'years of experience', value:"years_of_experience"},
            {text:'consultation fee', value:"consultation_fee"},
            {text:'email_id', value:"mail_id"},
            {text:'phone number', value:"phone_number"},
            

            
        ],
        doctors:[] 
        

    }),
    created(){
        console.log('inside created')
        this.initialize()
    },
    methods:{
        initialize()
        {this.loadDoctors()},



        loadDoctors(){

            this.doctors=[]
            doctorServices.getAlldoctors()
            doctorServices.getAlldoctors().then(item=>{
                    item.doctor.forEach(
                        doctor=>{this.doctors.push(doctor)}
                    )
                    
                

            });
            

        }
    }
  }
</script>