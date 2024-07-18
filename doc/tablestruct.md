user:
- primary:
- - OpenID: string
- column:
- - LoginTime: int
- - ExpiresIn: int
- - SystemToken: string
- - permission: int
doctor:
- primary:
- - OpenID: string
- column:
- - doctorAvatar: string
- - doctorName: string
- - doctorProfile: string
reserve:
- primary:
- - ID: int
- column:
- - DoctorID: string
- - Mobile: int
- - Name: string
- - Token: string
- - Time: int
- - Status: int