import Baseline from  '@/components/Baseline.vue';
import Doctor from '@/components/Doctor.vue'
import patient from '@/components/patient.vue'
import FullCalendar  from '@/components/calendar.vue'


const routes = [
    { name:'Baseline', path: '/', component: Baseline},
    { name:'Doctor', path: '/Doctor', component: Doctor},
    { name:'patient', path: '/patient', component:patient},
    {name:'FullCalendar ', path: '/calendar', component:FullCalendar},
   
];

export default routes;
 