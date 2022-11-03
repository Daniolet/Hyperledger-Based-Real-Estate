import { useEffect, useState } from "react"
import { FiLogOut } from "react-icons/fi"
import { FaUserNinja } from "react-icons/fa"
import { useRouter } from "next/router"
const Realestate = () => {
    const router = useRouter()
    const [acc, setacc] = useState<{ accountId: string; balance: number; userName: string }>()
    const [check, setcheck] = useState<boolean>(true)
    const [model, setmodel] = useState<boolean>(false)

    useEffect(() => {
        const ans = localStorage.getItem("account")
        if (ans == null) {
            router.push("/")
        }
        setacc(JSON.parse(ans as string))
    }, [])
    useEffect(() => {

    }, [check])


    return (
        <div className={`flex flex-col dark:bg-gray-900 overflow-x-hidden overflow-y-scroll h-screen w-screen items-center ${check ? "" : "main"}`}>

            <div className="navbar sticky top-0 z-50 text-white flex justify-between px-10 items-center w-full dark:bg-gray-700 dark:border-gray-700 dark:text-white py-4">
                <div className="flex space-x-6">
                    <button className="glass-button">Sellings</button>
                    <button className="glass-button">Donations</button>
                </div>
                <div className="font-semibold text-xl">
                    Block Chain Based Real Estate
                </div>
                <div className="flex items-center space-x-11">
                    <input type="checkbox" defaultChecked onClick={() => {
                        setcheck(!check)
                        localStorage.setItem("dark", check.toString())
                    }} className="hidden" id="darkmode-toggle" />
                    <label className={` ${check ? "" : "avoiddark"} swtch w-16 h-7 `} htmlFor="darkmode-toggle" />
                    <div className="font-bold p-5 hover:bg-slate-400 hover:text-black bg-slate-500 rounded-lg cursor-pointer" onClick={() => {
                        localStorage.removeItem("account")
                        router.push("/")
                    }}>
                        <FiLogOut className="text-xl " />
                    </div>
                </div>
            </div>
            <div className=" flex relative flex-col text-white justify-center items-center w-screen h-screen">
                <div>The Listed Properties are lsited here</div>
            </div>
            <button
                type='button'
                className='fixed bottom-7 right-11 inline-block  rounded-full bg-green-500 p-3 text-xs font-medium uppercase leading-tight text-white shadow-md transition duration-150 ease-in-out hover:bg-green-600 hover:shadow-lg focus:bg-green-400 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-green-800 active:shadow-lg '
                onClick={() => setmodel(!model)}
            >
                {/* <TbLetterM className='h-8 w-8' /> */}
                <FaUserNinja className='h-8 w-8' />
            </button>
            <div className={`fixed z-10 overflow-y-auto top-0 w-full left-0 ${model ? "" : "hidden"}`} id="modal">
                <div className="flex items-center justify-center min-height-100vh pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                    <div className="fixed inset-0 transition-opacity">
                        <div className="absolute inset-0 bg-gray-900 opacity-75" />
                    </div>
                    <span className="hidden sm:inline-block sm:align-middle sm:h-screen">&#8203;</span>
                    <div className="inline-block align-center bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full" role="dialog" aria-modal="true" aria-labelledby="modal-headline">
                        <div className="bg-white px-4 text-lg font-semibold text-slate-600 pt-5 pb-4 sm:p-6 sm:pb-4">
                            <label>Account Details</label>
                            <div className="px-3 py-4 bg-gray-200 rounded-xl m-5">
                                <div className="flex">
                                    <div>Account ID :</div>
                                    <div>{acc?.accountId}</div>
                                </div>
                                <div className="flex">
                                    <div>Balence Held :</div>
                                    <div>{acc?.balance}</div>
                                </div>
                                <div className="flex">
                                    <div>UserName :</div>
                                    <div>{acc?.userName}</div>
                                </div>
                            </div>
                            <div>
                                <i>
                                    After the sales, donation or pledge operations, the guarantee status is true

                                    When the guarantee status is false, the sale, donation or pledge operation can be initiated
                                </i>
                            </div>
                        </div>
                        <div className="bg-gray-200 px-4 py-3 w-full flex justify-between">
                            <a href="https://github.com/jayendramadaram/Hyperledger-Based-real-estate">
                                <button type="button" className="py-2 px-4 bg-green-500 text-white rounded hover:bg-gray-700 mr-2"><i className="fas fa-times"></i> VISIT PROJECT</button>
                            </a>
                            <button type="button" className="py-2 px-4 bg-gray-500 text-white rounded hover:bg-gray-700 mr-2" onClick={() => setmodel(!model)} ><i className="fas fa-times"></i> Cancel</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Realestate
