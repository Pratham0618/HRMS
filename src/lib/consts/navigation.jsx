import { FcGrid } from "react-icons/fc";
import { FcManager } from "react-icons/fc";
import { FcLibrary } from "react-icons/fc";
import { FcLeave } from "react-icons/fc";
import { FcFlowChart } from "react-icons/fc";
import { FcSettings } from "react-icons/fc";
import { FcExternal } from "react-icons/fc";

export const Sidebar_Links = [
    {
        key:'dashboard',
        label:'Dashboard',
        path:'/dashboard',
        icon:<FcGrid />
    },
    {
        key:'employee section',
        label:'Employee Section',
        path:'/employeelist',
        icon:<FcManager />
    },
    {
        key:'department section',
        label:'Department Section',
        path:'/departmentsection',
        icon:<FcLibrary />
    },
    {
        key:'leave types',
        label:'Leave Types',
        path:'/leavetypes',
        icon:<FcFlowChart />
    },
    {
        key:'leavemanagement',
        label:'Leave Management',
        path:'/leavemanagement',
        icon:<FcLeave />
    }
]

export const Sidebar_Links_Bottom = [
    {
        key:'settings',
        label:<div className="text-xl">Settings</div>,
        path:'/settings',
        icon:<FcSettings />
    },
    {
        key:'logout',
        label:<div className='text-red-400'>Logout</div>,
        path:'/',
        icon:<FcExternal />,
    }
]