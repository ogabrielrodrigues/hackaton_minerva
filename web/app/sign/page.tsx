import { SignSection } from '@/components/sections/sign_section'

export default async function SignPage() {
  return (
    <main className="h-[calc(100%-86px)] flex flex-col justify-center gap-4 px-5 md:w-1/2 md:mx-auto lg:w-2/5 xl:w-1/5">
      <SignSection />
    </main>
  )
}
