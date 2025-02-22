import React from "react";
import { useState , useEffect } from "react";

export default function Search()
{
    return (
        <div className="min-h-screen w-screen bg-[#0D1117] flex-col justify-center items-center">
            <div className="text-6xl font-extrabold text-white text-center font-['Roboto']">
                GitFindr
            </div>
            <div>
            
            <div class="grid bg-gradient-to-r from-[#0f0f10] via-transparent to-transparent bg-repeat-x bg-repeat-y absolute inset-0 z-[-1] filter blur-sm"></div>
<div id="poda" class="relative">
  <div class="glow absolute inset-0 z-[-1] max-h-[130px] max-w-[354px] overflow-hidden filter blur-[30px] opacity-40">
    <div class="absolute inset-0 transform translate-x-[-50%] translate-y-[-50%] rotate-[60deg] bg-conic-gradient from-[#000] via-[#402fb5] to-[#cf30aa]"></div>
  </div>
  
  <div class="darkBorderBg absolute inset-0 z-[-1] max-h-[65px] max-w-[312px] rounded-lg filter blur-sm">
    <div class="absolute inset-0 transform translate-x-[-50%] translate-y-[-50%] rotate-[82deg] bg-conic-gradient from-transparent via-[#18116a] to-[#6e1b60] filter brightness-125"></div>
  </div>

  <div class="white absolute inset-0 z-[-1] max-h-[63px] max-w-[307px] rounded-lg filter blur-sm">
    <div class="absolute inset-0 transform translate-x-[-50%] translate-y-[-50%] rotate-[83deg] bg-conic-gradient from-transparent via-[#a099d8] to-[#dfa2da] filter brightness-125"></div>
  </div>

  <div class="border absolute inset-0 z-[-1] max-h-[59px] max-w-[303px] rounded-lg filter blur-sm">
    <div class="absolute inset-0 transform translate-x-[-50%] translate-y-[-50%] rotate-[70deg] bg-conic-gradient from-[#1c191c] via-[#402fb5] to-[#cf30aa] filter brightness-125"></div>
  </div>

  <div id="main" class="relative">
    <input type="text" name="text" placeholder="Search..." class="input w-[301px] h-[56px] rounded-lg text-white bg-[#010201] px-[59px] text-lg focus:outline-none" />
    <div id="input-mask" class="absolute left-[70px] top-[18px] w-[100px] h-[20px] bg-gradient-to-r from-transparent to-black filter blur-sm pointer-events-none"></div>
    <div id="pink-mask" class="absolute left-[5px] top-[10px] w-[30px] h-[20px] bg-[#cf30aa] opacity-80 filter blur-[20px] transition-all duration-2000"></div>
    
    <div class="filterBorder absolute top-[7px] right-[7px] w-[40px] h-[42px] overflow-hidden rounded-lg">
      <div class="absolute inset-0 transform translate-x-[-50%] translate-y-[-50%] rotate-[90deg] bg-conic-gradient from-[#3d3a4f] via-transparent to-[#3d3a4f] filter brightness-[1.35]"></div>
    </div>

    <div id="filter-icon" class="absolute top-[8px] right-[8px] w-[38px] h-[40px] flex items-center justify-center z-[2] bg-gradient-to-b from-[#161329] to-[#1d1b4b] rounded-lg">
      <svg height="27" width="27" viewBox="4.8 4.56 14.832 15.408" fill="none" stroke="#d6d6e6" stroke-width="1" stroke-linecap="round" stroke-linejoin="round">
        <path d="M8.16 6.65002H15.83C16.47 6.65002 16.99 7.17002 16.99 7.81002V9.09002C16.99 9.56002 16.7 10.14 16.41 10.43L13.91 12.64C13.56 12.93 13.33 13.51 13.33 13.98V16.48C13.33 16.83 13.1 17.29 12.81 17.47L12 17.98C11.24 18.45 10.2 17.92 10.2 16.99V13.91C10.2 13.5 9.97 12.98 9.73 12.69L7.52 10.36C7.23 10.08 7 9.55002 7 9.20002V7.87002C7 7.17002 7.52 6.65002 8.16 6.65002Z"></path>
      </svg>
    </div>

    <div id="search-icon" class="absolute left-[20px] top-[15px]">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle stroke="url(#search)" r="8" cy="11" cx="11"></circle>
        <line stroke="url(#searchl)" y2="16.65" y1="22" x2="16.65" x1="22"></line>
        <defs>
          <linearGradient gradientTransform="rotate(50)" id="search">
            <stop stop-color="#f8e7f8" offset="0%"></stop>
            <stop stop-color="#b6a9b7" offset="50%"></stop>
          </linearGradient>
          <linearGradient id="searchl">
            <stop stop-color="#b6a9b7" offset="0%"></stop>
            <stop stop-color="#837484" offset="50%"></stop>
          </linearGradient>
        </defs>
      </svg>
    </div>
  </div>
</div>

            </div>
            <div>
                <img></img>
            </div>

        </div>
    )
}