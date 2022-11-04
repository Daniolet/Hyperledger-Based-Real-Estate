import { useEffect, useState } from "react"
import { FiLogOut } from "react-icons/fi"
import { FaUserNinja } from "react-icons/fa"
import { AiOutlinePlusCircle } from "react-icons/ai"
import { useRouter } from "next/router"
import toast, { Toaster } from "react-hot-toast"

const Selling = () => {
    const router = useRouter()
    const [acc, setacc] = useState<{ accountId: string; balance: number; userName: string }>()
    const [model, setmodel] = useState<boolean>(false)
    const [check, setcheck] = useState<boolean>(true)
    const [error, seterror] = useState<string>("Couldnt not save")
    const [property, setproperty] = useState<QuerySelling>()

    const [RealEstate, setRealEstate] = useState<[QuerySelling] | null>(null)



    useEffect(() => {
        const ans = localStorage.getItem("account")
        if (ans == null) {
            router.push("/")
        }
        setacc(JSON.parse(ans as string))
        const accounts = () => {
            fetch("http://localhost:8000/api/v1/queryAccountList", {
                method: "post"
            })
                .then((resp) => resp.json()
                    .then((res: Queryacc) => {
                        console.log(res);
                        // setaccs(res.data)
                    })).catch((err) => alert(err))
        }
        accounts()
        QueryRealEstate()
    }, [])


    const logger = () => {
        console.log(RealEstate, "try");
    }



    const Cancel = (buyer: string, saleobj: string) => {
        toast.promise(
            fetch("http://localhost:8000/api/v1/updateSelling", {
                method: "post",
                body: JSON.stringify({
                    buyer: buyer,
                    objectOfSale: saleobj,
                    seller: acc?.accountId,
                    status: "cancelled"
                })
            }).then((res) => res.json().then((resp) => {
                console.log(resp)
                QueryRealEstate()
            }).catch((err) => seterror(JSON.stringify(err)))),
            {
                loading: 'Saving...',
                success: <b>Settings saved!</b>,
                error: <b>{error}</b>,
            }
        )
    }

    const ConfirmBUY = (buyer: string, saleobj: string) => {
        toast.promise(
            fetch("http://localhost:8000/api/v1/updateSelling", {
                method: "post",
                body: JSON.stringify({
                    buyer: buyer,
                    objectOfSale: saleobj,
                    seller: acc?.accountId,
                    status: "done"
                })
            }).then((res) => res.json().then((resp) => {
                console.log(resp)
                QueryRealEstate()
            }).catch((err) => seterror(JSON.stringify(err)))),
            {
                loading: 'Saving...',
                success: <b>Settings saved!</b>,
                error: <b>{error}</b>,
            }
        )
    }

    const CreateBuy = (seller: string, saleobj: string) => {
        toast.promise(
            fetch("http://localhost:8000/api/v1/createSellingByBuy", {
                method: "post",
                body: JSON.stringify({
                    buyer: acc?.accountId,
                    objectOfSale: saleobj,
                    seller: seller,
                })
            }).then((res) => res.json().then((resp) => {
                console.log(resp)
                QueryRealEstate()
            }).catch((err) => seterror(JSON.stringify(err)))),
            {
                loading: 'Saving...',
                success: <b>Settings saved!</b>,
                error: <b>{error}</b>,
            }
        )
        QueryRealEstate()
    }



    const QueryRealEstate = () => {
        fetch("http://localhost:8000/api/v1/querySellingList", {
            method: "post",
        }).then((res) => res.json().then((resp) => {
            setRealEstate(resp.data)
        }).catch((err) => console.log(err)))

    }


    return (
        <div className={`flex flex-col dark:bg-gray-900 overflow-x-hidden overflow-y-scroll h-screen w-screen items-center ${check ? "" : "main"}`}>

            <div className="navbar sticky top-0 z-50 text-white flex justify-between px-10 items-center w-full dark:bg-gray-700 dark:border-gray-700 dark:text-white py-4">
                <div className="flex space-x-6">
                    <button className="glass-button" onClick={() => router.push("/realestate")}>Real Estate</button>
                    <button className="glass-button" onClick={() => router.push("/donation")}>Donations</button>
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
            <div className="w-full text-slate-200 text-2xl font-thin underline font-mono pt-10 px-8">
                Your Organization Products are listed Here
            </div>
            <div className=" flex relative flex-col text-white  items-center w-screen h-screen">
                <div >
                    {
                        RealEstate == null ? <div>The Listed Properties are Shown here</div> : (
                            <div className="grid grid-cols-3 gap-16 p-12">
                                {
                                    RealEstate.map((ele, index) => <div className={``} key={index}>
                                        <div className={`w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700`}>
                                            <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
                                                <h1 className="text-lg italic font-thin leading-tight tracking-tight text-gray-700 md:text-2xl dark:text-white">
                                                    {ele.sellingStatus}
                                                </h1>
                                                <div className="space-y-4 md:space-y-6" >
                                                    <div className="flex justify-center w-full">
                                                        <div className='w-full'>
                                                            <div className="dropdown relative w-full space-y-2">
                                                                <button
                                                                    className=" dropdown-toggle  px-6 py-2.5   font-medium text-xs leading-tight uppercase  shadow-md   focus:shadow-lg focus:outline-none focus:ring-0   active:shadow-lg active:text-white transition duration-150 ease-in-out  items-center whitespace-nowrap  justify-center bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 space-y-7 dark:placeholder-gray-400 dark:text-white "
                                                                    type="button"
                                                                    id="dropdownMenuButton1"
                                                                >
                                                                    <div>Property ID : {ele.objectOfSale}</div>
                                                                    <div>Onwer    ID : {ele.seller}</div>
                                                                    <div>Price Bidding : {ele.price}</div>
                                                                    <div>Validity : {ele.salePeriod}</div>
                                                                    <div>Time Of Creation : {ele.createTime}</div>
                                                                </button>
                                                                <button
                                                                    className=" dropdown-toggle  px-6 py-2.5   font-medium text-xs leading-tight uppercase  shadow-md   focus:shadow-lg focus:outline-none focus:ring-0   active:shadow-lg active:text-white transition duration-150 ease-in-out  items-center whitespace-nowrap  justify-center bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 space-y-7 dark:placeholder-gray-400 dark:text-white "
                                                                    type="button"
                                                                    id="dropdownMenuButton1"
                                                                >
                                                                    BuyerId : {ele.buyer == "" ? "No Bidding" : ele.buyer}
                                                                </button>
                                                                <div className={`${acc?.accountId == "5feceb66ffc8" || ele.sellingStatus == "Cancelled" || ele.sellingStatus == "Finish" ? "hidden" : ""}  flex justify-end w-full space-x-7  py-4`}>
                                                                    <div className={`${acc?.accountId == ele.seller ? "bg-red-500" : "bg-blue-700"} ${(ele.buyer != "" && ele.seller != acc?.accountId) ? "hidden" : ""}  cursor-pointer px-7 py-3 rounded-xl`} onClick={() => {
                                                                        // setorder(acc?.accountId == ele.seller ? "cancel" : "buy")
                                                                        {
                                                                            acc?.accountId == ele.seller ? Cancel(ele.buyer, ele.objectOfSale) : CreateBuy(ele.seller, ele.objectOfSale)
                                                                        }
                                                                    }}>{acc?.accountId == ele.seller ? "Cancel" : "Buy"}</div>
                                                                    <div className={`${acc?.accountId == ele.seller && ele.buyer !== "" && ele.sellingStatus != "Finish" ? "" : "hidden"} bg-green-500  cursor-pointer px-7 py-3 rounded-xl`} onClick={() => {
                                                                        ConfirmBUY(ele.buyer, ele.objectOfSale)
                                                                    }}>Confirm Payment</div>
                                                                </div>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                    </div>)
                                }
                            </div>
                        )
                    }
                </div>
            </div>
            <button
                type='button'
                className='fixed bottom-7 right-11 inline-block  rounded-full bg-green-500 p-3 text-xs font-medium uppercase leading-tight text-white shadow-md transition duration-150 ease-in-out hover:bg-green-600 hover:shadow-lg focus:bg-green-400 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-green-800 active:shadow-lg '
                onClick={() => setmodel(!model)}
            >
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


            <Toaster
                position="bottom-left"
                reverseOrder={false}
            />
        </div >
    )
}

export default Selling