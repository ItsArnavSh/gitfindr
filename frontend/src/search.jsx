"use client"
import Form from "./form"
import patharImg from "./assets/background.png"

export default function Search() {
    return (
        <div className="min-h-screen w-full bg-[#0D1117] flex flex-col items-center justify-center px-4 py-8 space-y-8">
            <div className="text-4xl sm:text-5xl md:text-6xl font-extrabold text-white text-center font-['Roboto']">
                GitFindr
            </div>
            <div className="w-full max-w-md mx-auto">
                <Form />
            </div>
            <div className="w-full max-w-2xl md:max-w-3xl lg:max-w-4xl xl:max-w-5xl">
                <img
                    src={patharImg || "/placeholder.svg"}
                    alt="Pathar Image"
                    className="w-full h-auto object-contain"
                />
            </div>

        </div>
    )
}

