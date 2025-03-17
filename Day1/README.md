# Day 1

## Info - Dual booting or multi-booting
<pre>
- Boot loader system software utilities that is installed in Hard Disk - Master Boot Record(MBR)
- Master Boot Record is Sector 0, Byte 0 in the hard disk which is just 512 bytes
- Boot Loader software normally has to fit within the 512 bytes, it gets installed in MBR
- When the system is booted, after the BIOS POSt is completed, the BIOS instructs the CPU to run the Boot Loader
- The Boot Loader utility searches the hard disk looking for Operating System installed in it, if the number of OS is more than one, it prepares a menu and presents the menu for us to choose the OS we wish to boot into
- Only one OS can be active at any point of time
- Examples
  - LILO ( Linux Loader - legacy boot loader that was used in older Linux distributions )
  - GRUB - almost all Linux distributions
  - BootCamp ( to support dual boot with Mac OS-X or Windows )
</pre>  

## Info - Hypervisor
<pre>
- virtualization technology
- with virtualization, we can run many OS side by side on the same machine ( laptop, desktop, workstation or Server )
- many OS can be actively run at the same time on the same machine
- two types of Hypervisors
  - Type 1 
    - Bare Metal Hypervisors
    - Doesn't require a Host OS to install the Hypervisor software
    - Which is used in Servers & Workstations
    - examples
      - VMWare vSphere/vCenter
      - KVM 
  - Type 2
    - is used in laptops/desktops/workstations
    - requires Host OS ( Windows, Mac, Linux )
    - examples
      - VMWare Workstation ( Windows & Linux )
      - VMWare Fusion ( Mac OS-X )
      - Parallels ( Mac OS-X )
      - Oracle VirtualBox ( Windows, Linux & Mac )
      - Microsoft Hyper-V ( Windows )
- each Guest OS is installed in a Virtual Machine (VM)
- each Guest OS is a fully functional Operating System with their own hardware and dedicated OS Kernel, etc
- for each VM, we must allocated dedicated hardware resources, hence it is considered a heavy weight virtualization technology
  - CPU Cores
  - RAM
  - Storage
  - Virtual Network Card(s)
  - Virtual Graphics Cards(s)
</pre>

## Info - Hypervisor - High Level Architecture
![Hypevisor](HypervisorHighLevelArchitecture.png)

## Container Technology
<pre>
- Containerization - is an application virtualization technology
- light-weight virtualization technology
- each container, represents one application or application component
- containers are not Operating System
- containers are just application process not a OS
- containers don't even has their own dedicated hardwares or Operating System kernel
- all the containers that runs on a system shares the underlying HOST OS Kernel and Hardware resources
- contaners have some similartity to a virtual machine
  - just like VMs, containers are also allocated with Ip address
  - just like VMs, containers has its own file systems
  - just like VMs, containers has their own network stack ( 7 OSI Layers )
  - just like VMs, containers has their own port range ( 0 - 65535 )
- Linux Kernel container related features
  - Namespace and
    - helps isolating one container from other containers
  - Control Groups (CGroups)
    - helps applying resource quota restrictions for each container
      - how many cpu one container can use at the max
      - how much RAM one container can use at the max
</pre>

## Info - Container Runtime
<pre>
- is a low-level software that manages container images and containers
- is not so user-friendly, hence normally no end-users use it directly
- examples
  - runC
  - CRI-O 
</pre>

## Info - Container Engine
<pre>
- is a high-level software that manages container images and containers
- provides user-friendly command to manage images and containers, even without low-level linux kernel knowledge one can manages manages and container with the help of container engine
- container engines internally makes use of Container Runtime to manage images and containers
- examples
  - Docker is a Container Engine that depends on containerd which internally depends on runC Container Runtime
  - Podman is a Container Engine that depends on CRI-O container runtime
</pre>

## Info - Container Images
<pre>
- is a blueprint of a container
- it is similar to Windows.iso OS images, with single Windows12.iso we can install Windows 12 on multiple machines, same way with Container image we can creates multiple container instances
</pre>

## Info - Containers
<pre>
- is an instance of a container image
- container represents an application 
- all the tools and sofwares installed in the container image are available in every container that is created out of the container image
</pre>
