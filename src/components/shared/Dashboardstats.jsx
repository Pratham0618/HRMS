import React from "react";
import { FaBriefcase } from "react-icons/fa";
import { GrUserManager } from "react-icons/gr";
import { HiBuildingLibrary } from "react-icons/hi2";
import { AiOutlineLoading3Quarters } from "react-icons/ai";
import { TiTickOutline } from "react-icons/ti";
import { ImCross } from "react-icons/im";

function Dashboardstats(){
    return(
        <div className="w-full">
            <div className="flex gap-4 mb-5">
                <BoxWrapper>
                    <FaBriefcase className="text-5xl text-white mx-auto"/>
                    <div className="text-center mt-2">
                        <div className="text-white text-3xl">Available Leave Types</div>
                    </div>
                    <div className="text-white text-5xl mt-auto">0</div>
                </BoxWrapper>
                <BoxWrapper>
                    <GrUserManager className="text-5xl text-white mx-auto"/>
                    <div className="text-center mt-2">
                        <div className="text-white text-3xl">Employees</div>
                    </div>
                    <div className="text-white text-5xl mt-auto">0</div>
                </BoxWrapper>
                <BoxWrapper>
                    <HiBuildingLibrary className="text-5xl text-white mx-auto"/>
                    <div className="text-center mt-2">
                        <div className="text-white text-3xl">Available Departments</div>
                    </div>
                    <div className="text-white text-5xl mt-auto">0</div>
                </BoxWrapper>
            </div>
            <div className="flex gap-4">
                <BoxWrapper>
                    <AiOutlineLoading3Quarters className="text-5xl text-white mx-auto"/>
                    <div className="text-center mt-2">
                        <div className="text-white text-3xl">Pending Applications</div>
                    </div>
                    <div className="text-white text-5xl mt-auto">0</div>
                </BoxWrapper>
                <BoxWrapper>
                    <TiTickOutline className="text-5xl text-white mx-auto"/>
                    <div className="text-center mt-2">
                        <div className="text-white text-3xl">Approved Applications</div>
                    </div>
                    <div className="text-white text-5xl mt-auto">0</div>
                </BoxWrapper>
                <BoxWrapper>
                    <ImCross className="text-5xl text-white mx-auto"/>
                    <div className="text-center mt-2">
                        <div className="text-white text-3xl">Declined Applications</div>
                    </div>
                    <div className="text-white text-5xl mt-auto">0</div>
                </BoxWrapper>
            </div>
        </div>
    )
}

export default Dashboardstats

function BoxWrapper({children}){
    return (
        <div className="bg-slate-800 rounded-sm p-4 flex-1 border border-gray-200 flex flex-col items-center justify-between h-60 text-white">
            {children}
        </div>
    )
}
