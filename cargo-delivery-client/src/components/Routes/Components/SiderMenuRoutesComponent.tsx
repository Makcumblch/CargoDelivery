import ListRoutes from "./SiderRoutesComponents/ListRoutes"
import RouteMenu from "./SiderRoutesComponents/RouteMenu";

const SiderMenuRoutesComponent = () => {

    return (
        <div className="grid grid-rows-[1fr_220px] gap-2 h-full">
            <ListRoutes />
            <RouteMenu />
        </div>
    );
}

export default SiderMenuRoutesComponent