import Link from 'next/link'
import { GetEmployee } from '@/hooks/get_employee'
import { Feedbacks } from '@/components/feedbacks'
import { filterAction } from '@/actions/filter'
import { LoginForm } from '@/components/forms/login_form'
import { AdministratorFeedbacks } from '@/components/sections/administrator_feedbacks'
import { EmployeeFeedbacks } from '@/components/sections/employee_feedbacks'
import { LoginSection } from '@/components/sections/login_section'
import { FeedbacksSection } from '@/components/sections/feedbacks_section'

/**
 * TEST: Admin
 * email: admin@admin.minerva.com
 * password: admin1
 *
 * TEST: Employee
 * email: john@ti.minerva.com
 * password: john01
 * */

export default async function HomePage() {
  const { employee } = await GetEmployee()

  return (
    <main className="h-[calc(100%-86px)] flex flex-col justify-center gap-4 px-5 md:w-1/2 md:mx-auto lg:w-2/5 xl:w-1/5">
      {!employee ? <LoginSection /> : <FeedbacksSection />}
    </main>
  )
}
