import { GetEmployee } from '@/hooks/get_employee'
import { AdministratorFeedbacks } from './administrator_feedbacks'
import { EmployeeFeedbacks } from './employee_feedbacks'

export async function FeedbacksSection() {
  const { employee } = await GetEmployee()

  return (
    <div className="h-full mt-6 flex flex-col justify-center gap-4">
      <h1 className="text-4xl font-bold">
        Olá, <strong className="text-primary">{employee?.name}</strong>
      </h1>
      <div className="flex gap-2">
        <h3>
          Setor: <strong>{employee?.sector}</strong>
        </h3>
        <h3>
          Unidade: <strong>{employee?.unit}</strong>
        </h3>
      </div>

      {employee?.administrator ? (
        <AdministratorFeedbacks />
      ) : (
        <EmployeeFeedbacks />
      )}
    </div>
  )
}
