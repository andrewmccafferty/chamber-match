FROM microsoft/dotnet:2.0-sdk AS build-env
WORKDIR /app

# copy csproj and restore as distinct layers
COPY /src/ChamberMatch.Api/*.csproj ./
RUN dotnet restore

# copy everything else and build
COPY /src/ChamberMatch.Api/. ./
RUN dotnet publish -c Release -o out

# build runtime image
FROM microsoft/aspnetcore:2.0.0 
WORKDIR /app
COPY --from=build-env /app/out ./
ENTRYPOINT ["dotnet", "ChamberMatch.Api.dll"]