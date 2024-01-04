import Image from 'next/image'

import Logo from '@/assets/logo.svg'
import Link from 'next/link'
import { Logout } from './logout'

export function Header() {
  return (
    <header className="flex justify-between items-center w-screen px-5 py-4 border-b border-zinc-200">
      <Link href="/" replace>
        <Image src={Logo} className="w-32" alt="Minerva Logo" />
      </Link>
      <Logout />
    </header>
  )
}
