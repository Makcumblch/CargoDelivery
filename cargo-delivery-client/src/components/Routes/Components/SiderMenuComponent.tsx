import ListRoutes from "./SiderComponents/ListRoutes"
import RouteMenu from "./SiderComponents/RouteMenu";

const SiderMenuComponent = () => {

    return (
        <div className="grid grid-rows-[1fr_220px] gap-2 h-full">
            <ListRoutes />
            <RouteMenu />
        </div>
    );
}

export default SiderMenuComponent