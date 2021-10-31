# EasyCache
_caching made easy_

## Local Testing
1. Clone EasyCache repo
2. `cd EasyCache`
3. `go run .` (service will listen on localhost:1323 by default)
4. Test routes (/health, GET /cache, POST /cache)

## TODO List
- Implement DLL + Updating
- Add tests for existing functions
- Consider adding auth module
- Add env file for config variables
- Add Config Vars to HEALTH
- Enable resource updates on CONFIG vars (that can be changed)
- Start time tracking for all actions + logging these