import { LogOut } from 'lucide-react'
import { logoutAction } from '@/actions/logout'
import { GetToken } from '@/hooks/get_token'

export async function Logout() {
  const { token } = await GetToken()

  return (
    <div>
      <div className="flex items-center">
        {token && (
          <form action={logoutAction}>
            <button
              type="submit"
              className="p-2 text-primary font-medium hover:bg-zinc-300/40 rounded transition-colors"
            >
              <LogOut className="cursor-pointer text-primary w-5 h-5" />
            </button>
          </form>
        )}
      </div>
    </div>
  )
}
